package casa.account.v1;

import edu.umd.cs.findbugs.annotations.NonNull;
import io.micronaut.data.jdbc.annotation.JdbcRepository;
import io.micronaut.data.model.query.builder.sql.Dialect;
import io.micronaut.data.repository.CrudRepository;

import javax.validation.constraints.NotNull;
import java.util.Optional;

@JdbcRepository(dialect = Dialect.MYSQL)
public interface CasaAccountRepository extends CrudRepository<CasaAccountEntity, String> {
    Optional<CasaAccountEntity> findById(@NotNull @NonNull String id);
}

