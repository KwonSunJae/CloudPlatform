package cloudplatform.udcs.exception;

public class ExistMemberException extends RuntimeException {
    public ExistMemberException() {
    }

    public ExistMemberException(String message) {
        super(message);
    }
}
