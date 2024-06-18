package cloudplatform.udcs.repository;

import cloudplatform.udcs.domain.Member;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;

public interface MemberRepository extends JpaRepository<Member, Long> {

    boolean existsByUuidAndUserId(String uuid, String userId);

    Optional<Member> findByRefreshToken(String refreshToken);

    Optional<Member> findByUuidAndUserId(String uuid, String userId);
}
