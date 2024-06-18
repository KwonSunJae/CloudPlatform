package cloudplatform.udcs.domain;

public record TokenResponse(
		String accessToken,
		String refreshToken
) {
}
