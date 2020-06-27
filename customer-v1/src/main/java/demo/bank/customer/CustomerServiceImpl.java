package demo.bank.customer;

import demo.bank.*;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Random;

@GrpcService
public class CustomerServiceImpl extends CustomerServiceGrpc.CustomerServiceImplBase {

  private static final Logger logger = LoggerFactory.getLogger(CustomerServiceImpl.class);
  private Random rand = new Random();

  @Override
  public void getCustomer(GetCustomerRequest req, StreamObserver<Customer> responseObserver) {
    String login = req.getCustomerId();
    logger.info("Retrieving Customer details for {}", login);

    Customer customer = Customer.newBuilder()
            .setCustomerId(login)
            .setName("Dummy Customer")
            .build();

    // random delay to demo tracing
    try {
      Thread.sleep(rand.nextInt(10)*100L);
    } catch (InterruptedException e) {
    }

    responseObserver.onNext(customer);
    responseObserver.onCompleted();
  }
}
