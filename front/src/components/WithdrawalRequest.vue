<template>
  <div>
    <form class="form-signin">
      <h1 class="h1 mb-3 fw-normal text-center">Withdrawal request</h1>
      <div v-if="!loaded && !completed" key="loading">
        <h2 class="mt-3 fw-light text-center">Loading...</h2>
      </div>
      <div v-if="loaded && !completed" key="requesting">
        <div class="mb-3">
          <div class="input-group input-group-lg">
            <div class="input-group-text">{{ account.currencyCode }}</div>
            <input type="number" min="0" step="1" class="form-control" placeholder="0" v-model="dollars">
            <span class="input-group-text">.00</span>
            <div class="alert alert-danger mt-4" role="alert" v-if="errors !== null">
              {{ errors }}
            </div>
          </div>
          <hr class="mt-4 mb-3">
          <div class="align-content-start">
            <p class="mb-2 text-muted">From:</p>
            <div class="row">
              <div class="col col-sm-4 col-md-3">
                <div class="icon bg-primary rounded-circle icon text-center">
                  <svg xmlns="http://www.w3.org/2000/svg" height="48px" viewBox="0 0 24 24" width="48px" fill="#FFFFFF"><path d="M0 0h24v24H0V0z" fill="none"/><path d="M16 6l2.29 2.29-4.88 4.88-4-4L2 16.59 3.41 18l6-6 4 4 6.3-6.29L22 12V6h-6z"/></svg>
                </div>
              </div>
              <div class="col col-sm-8 col-md-9">
                <h4 class="fw-bold">{{ portfolio.name }}</h4>
                <p class="fw-light text-muted">{{ formattedBalance }} available</p>
              </div>
            </div>
          </div>
          <hr class="mt-0 mb-4">
        </div>
        <button class="w-100 btn btn-lg btn-primary" type="submit" :disabled="submitting" @click.prevent="submitRequest">Submit</button>
      </div>
      <div v-if="loaded && completed" key="completing">
        <div class="mb-3">
          <h5>Success, your order sheet has been created.</h5>
          <div class="card mt-4 mb-4">
            <div class="card-header">
              Order sheet #{{ orderSheet.id }}
            </div>
            <div class="card-body">
              <table class="table">
                <thead>
                <tr>
                  <th scope="col">#</th>
                  <th scope="col">Type</th>
                  <th scope="col">£</th>
                  <th scope="col">ISIN</th>
                </tr>
                </thead>
                <tbody>
                  <tr v-for="instruction in orderSheet.instructions" v-bind:key="instruction.id">
                    <td>{{ instruction.id }}</td>
                    <td>{{ instruction.instructionType }}</td>
                    <td>{{ decimalNumber(instruction.amount) }}</td>
                    <td>{{ instruction.isin }}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
        <button class="w-100 btn btn-lg btn-outline-secondary" type="submit" @click.prevent="reload">Restart</button>
      </div>
    </form>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import axios from 'axios';

axios.defaults.baseURL = 'http://localhost:8080/api';

@Options({
  name: 'WithdrawalRequest',
  props: {
    customerId: Number,
    accountId: Number,
    portfolioId: Number,
    language: String,
    currency: String,
  },
  data() {
    return {
      completed: false,
      errors: null,
      submitting: false,
      formatter: null,
      customer: null,
      account: null,
      portfolio: null,
      dollars: null,
      orderSheet: null,
    };
  },
  computed: {
    loaded() {
      return this.customer !== null
        && this.account !== null
        && this.portfolio !== null;
    },
    formattedBalance() {
      return this.loaded
        ? this.formatter.format(this.decimalNumber(this.account.balance))
        : null;
    },
    accountIdInt() {
      return parseInt(this.accountId, 10);
    },
    withdrawalInt() {
      return parseInt(this.withdrawal, 10);
    },
  },
  methods: {
    makeFormatter() {
      this.formatter = new Intl.NumberFormat(this.language, {
        style: 'currency',
        currency: this.currency,
      });
    },
    decimalNumber(num :number): string {
      return num.toString().replace(/\b(\d+)(\d{2})\b/g, '$1.$2');
    },
    addCents(num :string) {
      return num !== null ? parseInt(`${num}00`, 10) : null;
    },
    setCustomer(customerId :string) {
      axios.get(`/customers/${customerId}`)
        .then((response) => {
          this.customer = response.data;
        });
    },
    setAccount(customerId :string, accountId :string) {
      axios.get(`/customers/${customerId}/accounts/${accountId}`)
        .then((response) => {
          this.account = response.data;
        });
    },
    setPortfolio(portfolioId :string) {
      axios.get(`/portfolios/${portfolioId}`)
        .then((response) => {
          this.portfolio = response.data;
        });
    },
    submitRequest() {
      this.submitting = true;
      this.errors = null;
      axios.post(`/customers/${this.customerId}/accounts/${this.accountId}/withdrawal-requests`, {
        accountId: this.accountIdInt,
        amount: this.addCents(this.dollars),
      })
        .then((response) => {
          this.orderSheet = Object.prototype.hasOwnProperty.call(response.data, 'orderSheet') ? response.data.orderSheet : null;
          this.completed = true;
        })
        .catch((error) => {
          if (error.response) {
            this.errors = error.response.data;
          }
        })
        .finally(() => {
          this.submitting = false;
        });
    },
    reload() {
      window.location.reload();
    },
  },
  mounted() {
    this.makeFormatter();
    this.setCustomer(this.customerId);
    this.setAccount(this.customerId, this.accountId);
    this.setPortfolio(this.portfolioId);
  },
})
export default class WithdrawalRequest extends Vue {
  customerId!: number

  accountId!: number

  portfolioId!: number

  language!: string

  currency!: string
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

.form-signin {
  width: 100%;
  max-width: 360px;
  padding: 15px;
  margin: auto;
}
.slide-fade-enter-active {
  transition: all 0.3s ease-out;
}

.slide-fade-leave-active {
  transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(20px);
  opacity: 0;
}

.alert {
  width:100%;
}

</style>
