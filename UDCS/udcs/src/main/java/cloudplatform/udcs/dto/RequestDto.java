package cloudplatform.udcs.dto;

import lombok.Data;

@Data
public class RequestDto {
    private String dest;
    private String method;
    private String data;
}
