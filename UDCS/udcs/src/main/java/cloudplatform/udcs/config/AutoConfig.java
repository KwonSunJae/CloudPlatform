package cloudplatform.udcs.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.client.HttpComponentsClientHttpRequestFactory;
import org.springframework.http.converter.StringHttpMessageConverter;
import org.springframework.web.client.RestTemplate;

import java.nio.charset.StandardCharsets;

@Configuration
public class AutoConfig {

    @Bean
    public RestTemplate restTemplate() {
//        HttpComponentsClientHttpRequestFactory factory
//                = new HttpComponentsClientHttpRequestFactory();
//
//        factory.setConnectTimeout(5000);
//        factory.setReadTimeout(5000);
//        factory.setBufferRequestBody(false);

//        RestTemplate restTemplate = new RestTemplate(factory);
//        restTemplate.getMessageConverters()
//                .add(0, new StringHttpMessageConverter(StandardCharsets.UTF_8));
        return new RestTemplate();
    }
}
