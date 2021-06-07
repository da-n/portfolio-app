<template>
  <div>
    <form class="form-signin">
      <h1 class="h1 mb-3 fw-normal text-center">Withdrawal request</h1>
      <div v-if="!loaded">
        <h2 class="mt-3 fw-light text-center">Loading...</h2>
      </div>
      <transition name="fade">
        <div v-if="loaded">
          <div class="mb-3">
            <div class="input-group input-group-lg">
              <div class="input-group-text">{{ account.currencyCode }}</div>
              <input type="number" min="0" step="1" class="form-control" id="autoSizingInputGroup" placeholder="0" v-model="withdrawal">
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
          <button class="w-100 btn btn-lg btn-primary" type="submit" @click.prevent="submitRequest">Submit</button>
        </div>
      </transition>
    </form>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import axios from 'axios';

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
      formatter: null,
      customer: null,
      account: null,
      portfolio: null,
      withdrawal: null,
    };
  },
  computed: {
    // basic check that everything has been fetched and set from API
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
  },
  methods: {
    /**
     * makeFormatter instantiates a currency formatter
     */
    makeFormatter() {
      this.formatter = new Intl.NumberFormat(this.language, {
        style: 'currency',
        currency: this.currency,
      });
    },
    /**
     * decimalNumber convert a number into a string decimal representation
     */
    decimalNumber(num :number): string {
      return num.toString().replace(/\b(\d+)(\d{2})\b/g, '$1.$2');
    },
    /**
     * setCustomer fetches and sets a customer from the API
     * @param customerId
     */
    setCustomer(customerId :string) {
      axios.get(`http://localhost:8080/api/customers/${customerId}`)
        .then((response) => {
          this.customer = response.data;
        });
    },
    /**
     * setAccount fetches and sets a customer from the API
     * @param customerId
     * @param accountId
     */
    setAccount(customerId :string, accountId :string) {
      axios.get(`http://localhost:8080/api/customers/${customerId}/accounts/${accountId}`)
        .then((response) => {
          this.account = response.data;
        });
    },
    /**
     * setPortfolio fetches and sets a customer from the API
     * @param portfolioId
     */
    setPortfolio(portfolioId :string) {
      axios.get(`http://localhost:8080/api/portfolios/${portfolioId}`)
        .then((response) => {
          this.portfolio = response.data;
        });
    },
    submitRequest() {
      console.log({
        accountId: this.accountId,
        amount: this.withdrawal,
      });
      axios.post(`http://localhost:8080/api/customers/${this.customerId}/accounts/${this.accountId}/withdrawal-requests`, {
        accountId: this.accountId,
        amount: this.withdrawal,
      })
        .then((response) => {
          console.log(response.data);
        });
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
  max-width: 330px;
  padding: 15px;
  margin: auto;
}

</style>
