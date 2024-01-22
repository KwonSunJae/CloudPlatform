package cloudplatform.udcs.domain;

import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor
public enum MemberAuthority {
	USER("ROLE_USER"),
	ADMIN("ROLE_ADMIN");

	private final String role;
}
