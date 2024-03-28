package cloudplatform.udcs.repository;

import cloudplatform.udcs.domain.Member;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;

public interface MemberRepository extends JpaRepository<Member, Long> {

    boolean existsBySsidAndUserId(String ssid, String userId);

    Optional<Member> findBySsidAndUserId(String ssid, String userId);
}
