package cloudplatform.udcs.controller;

import cloudplatform.udcs.dto.RequestDto;
import cloudplatform.udcs.service.TransactionService;
import cloudplatform.udcs.util.AuthenticationUtil;
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
    public ResponseEntity<Void> somethingToDo(
            @RequestBody RequestDto requestDto,
            @RequestHeader(value = "Access-Token") String access_token
    ) {


        transactionService.requestSOMS(requestDto, getMemberId());
        return ResponseEntity.noContent().build();
    }


}
