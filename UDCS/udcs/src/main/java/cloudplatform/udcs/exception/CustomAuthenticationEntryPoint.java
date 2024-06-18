package cloudplatform.udcs.exception;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.web.AuthenticationEntryPoint;

import java.io.IOException;
import java.io.PrintWriter;

import static org.springframework.http.MediaType.APPLICATION_JSON;

public class CustomAuthenticationEntryPoint implements AuthenticationEntryPoint {

	@Override
	public void commence(HttpServletRequest request, HttpServletResponse response,
			AuthenticationException authException) throws IOException {

		response.setCharacterEncoding("UTF-8");
		response.setContentType(APPLICATION_JSON.toString());
		response.setStatus(HttpServletResponse.SC_UNAUTHORIZED);

		response.setCharacterEncoding("UTF-8");
		response.setContentType(APPLICATION_JSON.toString());
		response.setStatus(HttpServletResponse.SC_UNAUTHORIZED);

		ObjectMapper objectMapper = new ObjectMapper();
		objectMapper.registerModule(new JavaTimeModule());
		objectMapper.disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS);

		PrintWriter writer = response.getWriter();
		writer.write(objectMapper.writeValueAsString("사용자 인증에 실패하였습니다."));
		writer.flush();
		writer.close();
	}
}

