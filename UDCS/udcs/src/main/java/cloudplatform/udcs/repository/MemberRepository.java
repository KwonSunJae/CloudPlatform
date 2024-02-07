package cloudplatform.udcs.repository;

import cloudplatform.udcs.domain.Member;
import org.springframework.data.jpa.repository.JpaRepository;

public interface MemberRepository extends JpaRepository<Member, Long> {

    boolean existsBySsidAndUserId(String ssid, String userId);

    Member findBySsidAndUserId(String ssid, String userId);
}
