package cloudplatform.udcs.service;

import cloudplatform.udcs.domain.Member;
import cloudplatform.udcs.domain.TransactionType;
import cloudplatform.udcs.dto.RequestDto;
import cloudplatform.udcs.repository.MemberRepository;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.client.RestTemplate;

import java.time.LocalDateTime;
import java.util.UUID;

@Service
@Slf4j
@RequiredArgsConstructor
public class TransactionService {

    private final MemberRepository memberRepository;
    private final RestTemplate restTemplate;

    @Value("${go.server.url}")
    private String SOMS_URL;


    @Transactional(propagation = Propagation.REQUIRED, rollbackFor = Exception.class)
    public void requestSOMS(RequestDto requestDto, Long memberId) {
        // 고유한 트랜잭션 ID 생성
        String transactionId = UUID.randomUUID().toString();

        String response = apiServer(requestDto, memberId, transactionId);

        logEnd(transactionId, UUID.randomUUID().toString());
    }

    private String apiServer(RequestDto requestDto, Long memberId, String transactionId) {
        String response = "";
        Member member = getMemberById(memberId);

        final HttpHeaders httpHeaders = new HttpHeaders();
        httpHeaders.set("ssid", member.getSsid());
        final HttpEntity<String> entity = new HttpEntity<>(httpHeaders);

        int maxAttempts = 2;
        int currentAttempt = 0;
        boolean success = false;

        logEnroll(requestDto, transactionId);

        while (currentAttempt < maxAttempts && !success) {

            if (currentAttempt != 0) {
                logTry(requestDto, transactionId, UUID.randomUUID().toString());
            }

            try {
                response = restTemplate.exchange(SOMS_URL + requestDto.getDest(), HttpMethod.valueOf(requestDto.getMethod()), entity, String.class).getBody();
                success = true;
            } catch (Exception e) {
                currentAttempt++;
            }

        }

        if (!success) {
            logError(transactionId, UUID.randomUUID().toString());
        }

        return response;
    }


    private void logEnroll(RequestDto requestDto, String transactionId) {
        log.info("Type = {}, Dest = {}, Method = {}, Data = {}, Time = {}, IP = {}, Transaction_Id = {}",
                TransactionType.ENROLL,
                requestDto.getDest(),
                requestDto.getMethod(),
                requestDto.getData(),
                LocalDateTime.now(),
                SOMS_URL,
                transactionId);
    }

    private void logEnd(String targetTransactionId,String transactionId) {
        log.info("Type = {}, Target_Transaction_Id = {}, Time = {}, Transaction_Id = {}",
                TransactionType.END,
                targetTransactionId,
                LocalDateTime.now(),
                transactionId);
    }

    private void logTry(RequestDto requestDto, String targetTransactionId, String transactionId) {
        log.info("Type = {}, Target_Transaction_Id = {}, Dest = {}, Method = {}, Data = {}, Time = {}, IP = {}, Transaction_Id = {}, ",
                TransactionType.TRY,
                targetTransactionId,
                requestDto.getDest(),
                requestDto.getMethod(),
                requestDto.getData(),
                LocalDateTime.now(),
                SOMS_URL,
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
