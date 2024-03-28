package cloudplatform.udcs.controller;

import cloudplatform.udcs.domain.Response;
import cloudplatform.udcs.domain.TokenResponse;
import cloudplatform.udcs.dto.RequestDto;
import cloudplatform.udcs.service.MemberService;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.annotation.security.PermitAll;
import jakarta.servlet.http.HttpServletRequest;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import static cloudplatform.udcs.util.AuthenticationUtil.getMemberId;

@Slf4j
@RestController
@RequiredArgsConstructor
@Tag(name = "AuthController", description = "회원가입 및 로그인 ,토큰 관리 Controller")
@RequestMapping("/auth")
public class MemberController {

    private final MemberService memberService;

    @PermitAll
    @PostMapping("/login")
    public ResponseEntity<TokenResponse> login(
            @RequestBody RequestDto requestDto,
            HttpServletRequest request
    ){
        return ResponseEntity.ok(memberService.login(requestDto,request.getRemoteAddr()));
    }

    @PermitAll
    @PostMapping("/signup")
    public ResponseEntity<Response> signUp(
            @RequestBody RequestDto requestDto,
            HttpServletRequest request
    ) {
        return ResponseEntity.ok(memberService.signUp(requestDto,request.getRemoteAddr()));
    }
}
