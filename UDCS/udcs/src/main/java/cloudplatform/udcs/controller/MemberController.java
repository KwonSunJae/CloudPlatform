package cloudplatform.udcs.controller;

import cloudplatform.udcs.domain.Response;
import cloudplatform.udcs.dto.RequestDto;
import cloudplatform.udcs.service.TransactionService;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.servlet.http.HttpServletRequest;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

import static cloudplatform.udcs.util.AuthenticationUtil.getMemberId;


@Slf4j
@RestController
@RequiredArgsConstructor
@Tag(name = "AuthController", description = "Member 관련 api")
@RequestMapping("/member")
public class MemberController {
    private final TransactionService transactionService;

    @PreAuthorize("hasAnyRole('USER')")
    @PostMapping("")
    public ResponseEntity<Response> getMemberInfo(
            @RequestBody RequestDto requestDto,
            @RequestHeader(value = "Access-Token") String access_token,
            HttpServletRequest request
    ) {
        return ResponseEntity.ok(transactionService.getMemberInfo(requestDto, getMemberId(),request.getRemoteAddr()));
    }
}
