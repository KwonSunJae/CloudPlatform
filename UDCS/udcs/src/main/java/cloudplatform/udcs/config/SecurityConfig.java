package cloudplatform.udcs.config;


import cloudplatform.udcs.exception.CustomAccessDeniedHandler;
import cloudplatform.udcs.exception.CustomAuthenticationEntryPoint;
import cloudplatform.udcs.jwt.Jwt;
import cloudplatform.udcs.jwt.JwtAuthenticationFilter;
import cloudplatform.udcs.jwt.JwtProperties;
import lombok.RequiredArgsConstructor;
import org.springframework.boot.autoconfigure.security.servlet.PathRequest;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityCustomizer;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.web.AuthenticationEntryPoint;
import org.springframework.security.web.SecurityFilterChain;
import org.springframework.security.web.access.AccessDeniedHandler;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;
import org.springframework.security.web.util.matcher.AntPathRequestMatcher;
import org.springframework.web.cors.CorsConfiguration;
import org.springframework.web.cors.CorsConfigurationSource;
import org.springframework.web.cors.UrlBasedCorsConfigurationSource;

import java.util.List;

@Configuration
@EnableWebSecurity
@RequiredArgsConstructor
public class SecurityConfig {

	private final JwtProperties jwtProperties;

	@Bean
	public Jwt jwt() {
		return new Jwt(
				jwtProperties.getClientSecret(),
				jwtProperties.getIssuer(),
				jwtProperties.getTokenExpire(),
				jwtProperties.getRefreshTokenExpire()
		);
	}

	@Bean
	public WebSecurityCustomizer configure() {
		return (web) -> web.ignoring().requestMatchers(
//				new AntPathRequestMatcher("/resources/**"),
                new AntPathRequestMatcher("/v3/api-docs/**"),
                new AntPathRequestMatcher("/swagger-resources/**"),
                new AntPathRequestMatcher("/swagger-ui/**"),
                new AntPathRequestMatcher("/webjars/**"),
                new AntPathRequestMatcher("/swagger/**"),
                new AntPathRequestMatcher("/swagger-ui.html"),
				PathRequest.toStaticResources().atCommonLocations());

	}

	@Bean
	CorsConfigurationSource corsConfigurationSource() {
		CorsConfiguration configuration = new CorsConfiguration();
		configuration.setAllowedOrigins(List.of("http://localhost:3000","http://117.16.136.172:10000"));
		configuration.setAllowedMethods(List.of("*"));
		configuration.setAllowedHeaders(List.of("*"));
//		configuration.setAllowedHeaders(
//				List.of(
//						"Origin", "Accept", "X-Requested-With", "Content-Type", "Access-Control-Request-Method",
//						"Access-Control-Request-Headers", "Authorization", "Access-Token", "Refresh-Token"
//				));
		configuration.setAllowCredentials(true);
		UrlBasedCorsConfigurationSource source = new UrlBasedCorsConfigurationSource();
		source.registerCorsConfiguration("/**", configuration);
		return source;
	}

	@Bean
	public AccessDeniedHandler accessDeniedHandler() {
		return new CustomAccessDeniedHandler();
	}

	@Bean
	public AuthenticationEntryPoint authenticationEntryPoint() {
		return new CustomAuthenticationEntryPoint();
	}

	@Bean
	public JwtAuthenticationFilter jwtAuthenticationFilter() {
		return new JwtAuthenticationFilter(jwt());
	}

	@Bean
	public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
		http.csrf((csrf) -> csrf.disable())

			.authorizeHttpRequests((authorizeHttpRequests) -> authorizeHttpRequests
							.requestMatchers("/**").permitAll()
							.anyRequest().authenticated()
			)
			.csrf((csrf) -> csrf.disable())
			.headers((headers) -> headers.disable())
			.formLogin((formLogin) -> formLogin.disable())
			.httpBasic((httpBasic) -> httpBasic.disable())
			.rememberMe((rememberMe) -> rememberMe.disable())
			.logout((logout) -> logout.disable())

			.sessionManagement((sessionManagement) ->
					sessionManagement.sessionCreationPolicy(SessionCreationPolicy.STATELESS))

			.exceptionHandling((exceptionHandling) -> exceptionHandling
					.authenticationEntryPoint(authenticationEntryPoint())
					.accessDeniedHandler(accessDeniedHandler()))

			.addFilterBefore(jwtAuthenticationFilter(), UsernamePasswordAuthenticationFilter.class)
			.cors((cors) -> corsConfigurationSource());

		return http.build();
	}
}
