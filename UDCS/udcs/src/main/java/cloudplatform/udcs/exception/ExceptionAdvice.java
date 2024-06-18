package cloudplatform.udcs.exception;

import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@RequiredArgsConstructor
@RestControllerAdvice
public class ExceptionAdvice{

    @ExceptionHandler(AuthException.class)
    private ResponseEntity<String> authException(Exception e) {
        return ResponseEntity.status(HttpStatusCode.valueOf(401)).body(e.getMessage());
    }

    @ExceptionHandler(ExistMemberException.class)
    private ResponseEntity<String> existMemberException(Exception e) {
        return ResponseEntity.status(HttpStatusCode.valueOf(409)).body(e.getMessage());
    }


}

