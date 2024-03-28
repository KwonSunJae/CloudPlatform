package cloudplatform.udcs.service;

import cloudplatform.udcs.domain.Member;
import cloudplatform.udcs.domain.Response;
import cloudplatform.udcs.domain.TransactionType;
import cloudplatform.udcs.dto.RequestDto;
import cloudplatform.udcs.repository.MemberRepository;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.client.HttpServerErrorException;
import org.springframework.web.client.RestTemplate;

import java.time.LocalDateTime;
import java.util.UUID;

@Service
@Slf4j
@RequiredArgsConstructor
public class TransactionService {

    private final MemberRepository memberRepository;
    private final RestTemplate restTemplate;

    private final ObjectMapper mapper;
    private final int maxAttempts = 2;


    @Value("${go.server.url}")
    private String SOMS_URL;


    @Transactional(propagation = Propagation.REQUIRED, rollbackFor = Exception.class)
    public Response requestSOMS(RequestDto requestDto, Long memberId, String remoteAddr) {
        // 고유한 트랜잭션 ID 생성
        String transactionId = UUID.randomUUID().toString();

        Response response = apiServer(requestDto, memberId, transactionId, remoteAddr);

        logEnd(transactionId, UUID.randomUUID().toString());
        return response;
    }

    private Response apiServer(RequestDto requestDto, Long memberId, String transactionId, String remoteAddr){
        Member member = getMemberById(memberId);
        HttpHeaders httpHeaders = new HttpHeaders();
        httpHeaders.set("ssid", member.getSsid());

        logEnroll(requestDto, transactionId, remoteAddr);

        if(requestDto.getMethod().equals(HttpMethod.GET.name())){
            return nonePolicy(requestDto, new HttpEntity<>(httpHeaders), transactionId, remoteAddr);
        }else{
            return tryPolicy(requestDto, new HttpEntity<>(requestDto.getData(), httpHeaders), transactionId, remoteAddr);
        }
    }

    private Response nonePolicy(RequestDto requestDto, HttpEntity<String> entity, String transactionId, String remoteAddr) {
        Response response = null;
        try{
            ResponseEntity<Response> responseEntity = restTemplate.exchange(SOMS_URL + requestDto.getDest(), HttpMethod.valueOf(requestDto.getMethod()), entity, Response.class);
            response = responseEntity.getBody();
        }catch (HttpServerErrorException e) {
            // 500 오류 발생 시
            try {
                response = mapper.readValue(e.getResponseBodyAsString(), Response.class);
            } catch (JsonProcessingException ex) {
                throw new RuntimeException(ex);
            }
        }
        return response;
    }

    private Response tryPolicy(RequestDto requestDto, HttpEntity<String> entity, String transactionId, String remoteAddr) {
        Response response = null;
        int currentAttempt = 0;
        boolean success = false;
        while (currentAttempt < maxAttempts && !success) {

            if (currentAttempt != 0) {
                logTry(requestDto, transactionId, UUID.randomUUID().toString(), remoteAddr);
            }

            try {
                ResponseEntity<Response> responseEntity = restTemplate.exchange(SOMS_URL + requestDto.getDest(), HttpMethod.valueOf(requestDto.getMethod()), entity, Response.class);

                if (responseEntity.getBody().getStatus() == 500) {
                    currentAttempt++;
                    continue;
                }

                response = responseEntity.getBody();
                success = true;
            } catch (HttpServerErrorException e) {
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

    private Member getMemberById(Long memberId) {
        return memberRepository.findById(memberId).orElseThrow(RuntimeException::new);
    }
}
