package com.example.model.msg;
import com.google.gson.Gson;
import lombok.Getter;

import java.nio.charset.StandardCharsets;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.Locale;
import java.util.Map;

@Getter
public class Message {
    private final int status;
    private final String message;
    private final String ts;
    private Map<String, Object> response;

    public Message(MessageStatus resultStatus) {
        this.status = resultStatus.getStatus();
        this.message = resultStatus
                .getDetail();
        this.ts = LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyyMMddHHmmss"));
        System.out.println("result : " + this.toString());
    }

    public Message(MessageStatus resultStatus, Map<String, Object> response, Locale locale) {
        this.status = resultStatus.getStatus();
        resultStatus.localized(locale);
        this.message = resultStatus.getDetail();
        this.ts = LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyyMMddHHmmss"));
        this.response = response;
        System.out.println("result : " + this.toString());
    }


//    public Message(int status, String message, Map<String, Object> response) {
//        this.status = status;
//        this.message = message;
//        this.ts = LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyyMMddHHmmss"));
//        this.response = response;
//        System.out.println("result : " + this.toString());
//    }

    @Override
    public String toString() {
        return new Gson().toJson(this);
    }

    public byte[] toByte() {
        return  new Gson().toJson(this).getBytes(StandardCharsets.UTF_8);
    }
}
