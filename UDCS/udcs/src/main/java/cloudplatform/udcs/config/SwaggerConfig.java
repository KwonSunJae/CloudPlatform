package cloudplatform.udcs.config;

import io.swagger.v3.oas.annotations.OpenAPIDefinition;
import io.swagger.v3.oas.annotations.info.Info;
import io.swagger.v3.oas.models.Components;
import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.servers.Server;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.util.Arrays;

@OpenAPIDefinition(
        info = @Info(title = "Test",
                description = "Test",
                version = "v1"))
@RequiredArgsConstructor
@Configuration
public class SwaggerConfig {

    @Value("${udcs.server.url}")
    private String UDCS_URL;


    @Bean
    public OpenAPI openAPI() {
        io.swagger.v3.oas.models.info.Info info = new io.swagger.v3.oas.models.info.Info()
                .title("API") // 타이틀
                .version("v1") // 문서 버전
                .description("잘못된 부분이나 오류 발생 시 바로 말씀해주세요."); // 문서 설명

        Server localServer = new Server();
        localServer.setDescription("HTTP");
        localServer.setUrl("http://localhost:8080");
        //
        Server testServer = new Server();
        testServer.setDescription("HTTP");
        testServer.setUrl(UDCS_URL);


        return new OpenAPI()
                .info(info)
                .components(new Components())
                .servers(Arrays.asList(testServer, localServer));

    }
}
