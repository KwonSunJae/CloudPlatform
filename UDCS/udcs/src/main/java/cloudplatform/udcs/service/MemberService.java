package cloudplatform.udcs.service;

import cloudplatform.udcs.domain.*;
import cloudplatform.udcs.dto.RequestDto;
import cloudplatform.udcs.exception.AuthException;
import cloudplatform.udcs.jwt.Jwt;
import cloudplatform.udcs.repository.MemberRepository;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.client.HttpServerErrorException;
import org.springframework.web.client.RestTemplate;

import java.time.LocalDateTime;
import java.util.Map;
import java.util.UUID;

@Service
@Slf4j
@RequiredArgsConstructor
public class MemberService {

    private final MemberRepository memberRepository;
    private final Jwt jwt;

    private final RestTemplate restTemplate;

    private final ObjectMapper mapper;

    @Value("${go.server.url}")
    private String SOMS_URL;

    @Transactional
    public Response signUp(RequestDto requestDto,String remoteAddr) {

        String transactionId = UUID.randomUUID().toString();

        Response response = apiServer(requestDto, transactionId, remoteAddr);
        String ssid = response.getData();
        String userId="";

        try {
            Map<String, String> userDataMap = mapper.readValue(requestDto.getData(), new TypeReference<Map<String, String>>() {});
            userId = userDataMap.get("UserID");
        } catch (Exception e) {
            e.printStackTrace();
        }

        if (!memberRepository.existsBySsidAndUserId(ssid, userId)) {
            memberRepository.save(Member.builder()
                    .userId(userId)
                    .ssid(ssid)
                    .memberAuthority(MemberAuthority.USER)
                    .build());
        }

        logEnd(transactionId, UUID.randomUUID().toString());
        return response;
    }


    @Transactional
    public TokenResponse login(RequestDto requestDto,String remoteAddr) {

        String transactionId = UUID.randomUUID().toString();

        Response response = apiServer(requestDto, transactionId, remoteAddr);
        String ssid = response.getData();
        String userId="";

        try {
            Map<String, String> userDataMap = mapper.readValue(requestDto.getData(), new TypeReference<Map<String, String>>() {});
            userId = userDataMap.get("UserID");
        } catch (Exception e) {
            e.printStackTrace();
        }

        logEnd(transactionId, UUID.randomUUID().toString());
        Member member = memberRepository.findBySsidAndUserId(ssid, userId)
                .orElseThrow(() -> new AuthException("해당 회원을 찾을 수 없습니다."));

        return publishToken(member);
    }

    @Transactional
    public TokenResponse publishToken(Member member)
    {
        TokenResponse tokenResponse = jwt.generateAllToken(
                Jwt.Claims.from(
                        member.getId(),
                        new String[] {
                                member.getMemberAuthority().getRole()
                        })
        );

        member.setRefreshToken(tokenResponse.refreshToken());

        return tokenResponse;
    }


    private Response apiServer(RequestDto requestDto, String transactionId, String remoteAddr){
        Response response = null;

        final HttpHeaders httpHeaders = new HttpHeaders();

        int maxAttempts = 2;
        int currentAttempt = 0;
        boolean success = false;

        logEnroll(requestDto, transactionId, remoteAddr);

        while (currentAttempt < maxAttempts && !success) {

            if (currentAttempt != 0) {
                logTry(requestDto, transactionId, UUID.randomUUID().toString(), remoteAddr);
            }

            try {
                HttpEntity<String> entity = new HttpEntity<>(httpHeaders);

                if(requestDto.getMethod().equals(HttpMethod.POST.name())){
                    entity = new HttpEntity<>(requestDto.getData(),httpHeaders);
                }

                ResponseEntity<Response> responseEntity = restTemplate.exchange(SOMS_URL + requestDto.getDest(), HttpMethod.valueOf(requestDto.getMethod()), entity, Response.class);

                if (responseEntity.getBody().getStatus() == 500) {
                    currentAttempt++;
                    continue;
                }
                if (responseEntity.getBody().getStatus() == 400) {
                    throw new AuthException("올바르지 않은 로그인 정보입니다.");
                }

                response = responseEntity.getBody();
                success = true;
            }
            catch (HttpServerErrorException e) {
                // 500 오류 발생 시
                try {
                    response = mapper.readValue(e.getResponseBodyAsString(), Response.class);
                } catch (JsonProcessingException ex) {
                    throw new RuntimeException(ex);
                }

                break;
            }
        }

        if (!success) {
            logError(transactionId, UUID.randomUUID().toString());
        }

        return response;
    }


    private void logEnroll(RequestDto requestDto, String transactionId,String remoteAddr) {
        log.info("Type = {}, Dest = {}, Method = {}, Data = {}, Time = {}, IP = {}, Transaction_Id = {}",
                TransactionType.ENROLL,
                requestDto.getDest(),
                requestDto.getMethod(),
                requestDto.getData(),
                LocalDateTime.now(),
                remoteAddr,
                transactionId);
    }

    private void logEnd(String targetTransactionId,String transactionId) {
        log.info("Type = {}, Target_Transaction_Id = {}, Time = {}, Transaction_Id = {}",
                TransactionType.END,
                targetTransactionId,
                LocalDateTime.now(),
                transactionId);
    }

    private void logTry(RequestDto requestDto, String targetTransactionId, String transactionId, String remoteAddr) {
        log.info("Type = {}, Target_Transaction_Id = {}, Dest = {}, Method = {}, Data = {}, Time = {}, IP = {}, Transaction_Id = {}, ",
                TransactionType.TRY,
                targetTransactionId,
                requestDto.getDest(),
                requestDto.getMethod(),
                requestDto.getData(),
                LocalDateTime.now(),
                remoteAddr,
                transactionId
        );
    }

    private void logError(String targetTransactionId, String transactionId) {
        log.info("Type = {}, Target_Transaction_Id = {}, Time = {}, Transaction_Id = {}",
                TransactionType.ERROR,
                targetTransactionId,
                LocalDateTime.now(),
                transactionId);
    }

}
