package demo.bank.customer;

import brave.Tracer;
import demo.bank.Customer;
import demo.bank.CustomerServiceGrpc;
import demo.bank.GetCustomerRequest;
import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpStatus;
import org.springframework.web.client.HttpClientErrorException;

import java.util.Map;
import java.util.Random;

@GrpcService
public class CustomerServiceImpl extends CustomerServiceGrpc.CustomerServiceImplBase {

  private static final Logger logger = LoggerFactory.getLogger(CustomerServiceImpl.class);
  private Random rand = new Random();
  private static final String RUNTIME =
      System.getProperty("java.runtime.name") + " " + System.getProperty("java.version");

  private Tracer tracer;
  private ExternalCustomerRepository repository;

  @Value("${spring.application.name}")
  private String appName;

  @Autowired
  public CustomerServiceImpl(Tracer tracer, ExternalCustomerRepository repository) {
    this.tracer = tracer;
    this.repository = repository;
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

    try {
      Map<String, String> extCustomer = repository.getExternalCustomerById(login);
      Customer customer =
          Customer.newBuilder()
              .setCustomerId(extCustomer.get("id"))
              .setName(extCustomer.get("name"))
              .setLoginName(extCustomer.get("login_name"))
              .build();
      responseObserver.onNext(customer);
      responseObserver.onCompleted();
      return;
    } catch (HttpClientErrorException e) {
      if (e.getStatusCode() == HttpStatus.NOT_FOUND) {
        responseObserver.onError(Status.NOT_FOUND.asException());
        return;
      }
    } catch (Exception e) {
      logger.warn("exception ", e);
    }
    responseObserver.onError(Status.INTERNAL.asException());
  }
}
