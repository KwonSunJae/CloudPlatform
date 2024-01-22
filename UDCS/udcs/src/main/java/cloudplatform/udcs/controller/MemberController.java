package cloudplatform.udcs.controller;

import cloudplatform.udcs.domain.TokenResponse;
import cloudplatform.udcs.service.MemberService;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.annotation.security.PermitAll;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@Slf4j
@RestController
@RequiredArgsConstructor
@Tag(name = "AuthController", description = "회원가입 및 로그인 ,토큰 관리 Controller")
@RequestMapping("/auth")
public class MemberController {

    private final MemberService memberService;

    @PermitAll
    @GetMapping("/test")
    public ResponseEntity<TokenResponse> testSignup(
            @RequestParam String ssid,
            @RequestParam String userId
    ){
        return ResponseEntity.ok(memberService.testSignup(ssid, userId));
    }
}
