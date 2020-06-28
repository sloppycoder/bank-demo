package demo.bank.customer;

import brave.Tracer;
import demo.bank.Customer;
import demo.bank.CustomerServiceGrpc;
import demo.bank.GetCustomerRequest;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;

import java.util.Random;

@GrpcService
public class CustomerServiceImpl extends CustomerServiceGrpc.CustomerServiceImplBase {

  private static final Logger logger = LoggerFactory.getLogger(CustomerServiceImpl.class);
  private Random rand = new Random();
  private static final String RUNTIME =
          System.getProperty("java.runtime.name") + " " + System.getProperty("java.version");

  Tracer tracer;

  @Value("${spring.application.name}")
  private String appName;

  @Autowired
  public CustomerServiceImpl(Tracer tracer) {
    this.tracer = tracer;
  }

  @Override
  public void getCustomer(GetCustomerRequest req, StreamObserver<Customer> responseObserver) {
    String login = req.getCustomerId();

    tracer
        .currentSpan()
        .tag("service", appName)
        .tag("runtime", RUNTIME)
        .tag(appName + "/get_customer/customer_id", login);

    logger.info("Retrieving Customer details for {}", login);

    Customer customer =
        Customer.newBuilder().setCustomerId(login).setName("Dummy Customer").build();

    // random delay to demo tracing
    try {
      Thread.sleep(rand.nextInt(10) * 100L);
    } catch (InterruptedException e) {
    }

    responseObserver.onNext(customer);
    responseObserver.onCompleted();
  }
}
