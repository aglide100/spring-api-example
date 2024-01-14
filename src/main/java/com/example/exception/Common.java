package com.example.exception;

import com.example.model.msg.MessageStatus;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public abstract class Common extends RuntimeException {
    protected int status;
    protected String message;

    protected Common(int status, String message) {
        super(message);
        this.status = status;
        this.message = message;
    }

    protected Common(int status, String message, Throwable cause) {
        super(message, cause);
        this.status = status;
        this.message = message;
    }

    protected Common(MessageStatus resultStatus) {
        super(resultStatus.getDetail());
        this.status = resultStatus.getStatus();
        this.message = resultStatus.getDetail();
    }

    protected Common(MessageStatus resultStatus, Throwable cause) {
        super(resultStatus.getDetail(), cause);
        this.status = resultStatus.getStatus();
        this.message = resultStatus.getDetail();
    }
}