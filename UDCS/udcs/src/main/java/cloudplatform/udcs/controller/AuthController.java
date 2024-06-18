package cloudplatform.udcs.controller;

import cloudplatform.udcs.domain.Response;
import cloudplatform.udcs.domain.TokenResponse;
import cloudplatform.udcs.dto.RequestDto;
import cloudplatform.udcs.service.MemberService;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.annotation.security.PermitAll;
import jakarta.servlet.http.Cookie;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.antlr.v4.runtime.Token;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import static cloudplatform.udcs.util.AuthenticationUtil.getMemberId;

@Slf4j
@RestController
@RequiredArgsConstructor
@Tag(name = "AuthController", description = "회원가입 및 로그인 ,토큰 관리 Controller")
@RequestMapping("/auth")
public class AuthController {

    private final MemberService memberService;

    @PermitAll
    @PostMapping("/login")
    public ResponseEntity<TokenResponse> login(
            @RequestBody RequestDto requestDto,
            HttpServletRequest request,
            HttpServletResponse response,
            @RequestParam(name = "isMemorized") boolean isMemorized
    ){
        TokenResponse tokenResponse = memberService.login(requestDto, request.getRemoteAddr());
        Cookie refreshTokenCookie = new Cookie("refreshToken", tokenResponse.refreshToken());
        refreshTokenCookie.setHttpOnly(true);
        refreshTokenCookie.setPath("/");
        refreshTokenCookie.setMaxAge(60 * 60 * 24 * 7); // 7일 -> 이거 변수로 refactoring 해야함

        refreshTokenCookie.setSecure(true);

        response.addCookie(refreshTokenCookie);
        //수정 필요
        return ResponseEntity.ok().body(tokenResponse);
    }

    @PermitAll
    @PostMapping("/signup")
    public ResponseEntity<Response> signUp(
            @RequestBody RequestDto requestDto,
            HttpServletRequest request
    ) {
        return ResponseEntity.ok(memberService.signUp(requestDto,request.getRemoteAddr()));
    }

    @PermitAll
    @PostMapping("/reissue")
    public ResponseEntity<TokenResponse> reissue(
            HttpServletRequest request, HttpServletResponse response
    ) {
        TokenResponse tokenResponse = memberService.reissue(
                request.getHeader("Authorization"),
                request.getRemoteAddr()
        );

        Cookie refreshTokenCookie = new Cookie("refreshToken", tokenResponse.refreshToken());

        refreshTokenCookie.setHttpOnly(true);
        refreshTokenCookie.setPath("/");
        refreshTokenCookie.setMaxAge(60 * 60 * 24 * 7);

        response.addCookie(refreshTokenCookie);

        return ResponseEntity.ok().body(tokenResponse);
    }
}
