package cloudplatform.udcs.domain;

import jakarta.persistence.*;
import lombok.*;

@Entity
@Getter
@NoArgsConstructor
@AllArgsConstructor
@Setter
@Builder
public class Member extends BaseEntity{

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(name = "member_authority", nullable = false)
    @Enumerated(EnumType.STRING)
    private MemberAuthority memberAuthority;

    private String ssid; //uuid
    private String userId; // user Id ex ehdtndla23

    private String refreshToken;

}
