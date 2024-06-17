package cloudplatform.udcs.domain;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

@Data
public class Response {
    @JsonProperty("data")
    private Object data;

    @JsonProperty("status")
    private int status;

    @JsonProperty("error")
    private Object error;

    @JsonIgnore
    public String getDataAsString() {
        if (data instanceof String) {
            return (String) data;
        }
        return null;
    }

    @JsonIgnore
    public String[] getDataAsStringArray() {
        if (data instanceof String[]) {
            return (String[]) data;
        }
        return null;
    }
}
