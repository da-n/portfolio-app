import { expect } from 'chai';
import { shallowMount } from '@vue/test-utils';
import WithdrawalRequest from '@/components/WithdrawalRequest.vue';

describe('WithdrawalRequest.vue', () => {
  it('renders props.msg when passed', () => {
    const msg = 'new message';
    const wrapper = shallowMount(WithdrawalRequest, {
      props: { msg },
    });
    expect(wrapper.text()).to.include(msg);
  });
});
