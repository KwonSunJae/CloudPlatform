package cloudplatform.udcs.domain;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

@Data
public class Response {
    @JsonProperty("data")
    private String data;

    @JsonProperty("status")
    private int status;

    @JsonProperty("error")
    private Object error;
}
