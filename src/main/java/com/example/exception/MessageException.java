package com.example.exception;

import com.example.model.msg.MessageStatus;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public abstract class MessageException extends RuntimeException {
    protected int status;
    protected String message;

    protected MessageException(int status, String message) {
        super(message);
        this.status = status;
        this.message = message;
    }

    protected MessageException(int status, String message, Throwable cause) {
        super(message, cause);
        this.status = status;
        this.message = message;
    }

    public MessageException(MessageStatus localized) {
        this.message = localized.getDetail();
        this.status = localized.getStatus();
    }
}