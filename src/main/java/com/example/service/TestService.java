package com.example.service;

import com.example.pb.svc.test.GreetingReq;
import com.example.pb.svc.test.GreetingRes;
import com.example.pb.svc.test.TestGrpc;
import net.devh.boot.grpc.client.inject.GrpcClient;
import org.springframework.stereotype.Service;

@Service
public class TestService {

    @GrpcClient("test")
    private TestGrpc.TestBlockingStub testStub;

    public String Testing(final String some) {
        try {
            GreetingRes response = this.testStub.greeting(GreetingReq.newBuilder().setSome(some).build());

            return response.toString();
        } catch (Exception e) {
            System.out.println("error : " + e.getMessage());
            return "";
        }
    }
}
