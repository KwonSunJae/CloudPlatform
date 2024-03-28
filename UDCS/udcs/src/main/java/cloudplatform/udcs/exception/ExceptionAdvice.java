package cloudplatform.udcs.exception;

import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@RequiredArgsConstructor
@RestControllerAdvice
public class ExceptionAdvice{

    @ExceptionHandler(AuthException.class)
    private ResponseEntity<String> authException(Exception e) {
        return ResponseEntity.internalServerError().body(e.getMessage());
    }

}

