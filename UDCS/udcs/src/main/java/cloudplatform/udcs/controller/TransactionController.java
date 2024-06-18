package cloudplatform.udcs.controller;

import cloudplatform.udcs.domain.Response;
import cloudplatform.udcs.dto.RequestDto;
import cloudplatform.udcs.service.TransactionService;
import cloudplatform.udcs.util.AuthenticationUtil;
import jakarta.servlet.http.HttpServletRequest;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RestController;

import static cloudplatform.udcs.util.AuthenticationUtil.*;

@RestController
@RequiredArgsConstructor
public class TransactionController {

    private final TransactionService transactionService;



    @PreAuthorize("hasAnyRole('USER')")
    @PostMapping("/transaction")
    public ResponseEntity<Response> somethingToDo(
            @RequestBody RequestDto requestDto,
            @RequestHeader(value = "Access-Token") String access_token,
            HttpServletRequest request
    ) {
        return ResponseEntity.ok(transactionService.requestSOMS(requestDto, getMemberId(),request.getRemoteAddr()));
    }

}
