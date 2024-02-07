package cloudplatform.udcs.service;

import cloudplatform.udcs.domain.Member;
import cloudplatform.udcs.domain.MemberAuthority;
import cloudplatform.udcs.domain.TokenResponse;
import cloudplatform.udcs.jwt.Jwt;
import cloudplatform.udcs.repository.MemberRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
@RequiredArgsConstructor
public class MemberService {

    private final MemberRepository memberRepository;
    private final Jwt jwt;

    @Transactional
    public TokenResponse testSignup(String ssid, String userId) {

        if (!memberRepository.existsBySsidAndUserId(ssid, userId)) {
            memberRepository.save(Member.builder()
                    .userId(userId)
                    .ssid(ssid)
                    .memberAuthority(MemberAuthority.USER)
                    .build());
        }

        Member member = memberRepository.findBySsidAndUserId(ssid, userId);

        return publishToken(member);
    }

    @Transactional
    public TokenResponse publishToken(Member member)
    {
        TokenResponse tokenResponse = jwt.generateAllToken(
                Jwt.Claims.from(
                        member.getId(),
                        new String[] {
                                member.getMemberAuthority().getRole()
                        })
        );

        member.setRefreshToken(tokenResponse.refreshToken());

        return tokenResponse;
    }
}
