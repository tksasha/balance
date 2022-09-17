# frozen_string_literal: true

RSpec.describe BalanceHelper, type: :helper do
  subject { helper }

  describe '#balance' do
    let(:params) { { currency: 'usd' } }

    before { allow(subject).to receive(:params).and_return(params) }

    before { allow(CalculateBalanceService).to receive(:call).with('usd').and_return(22.15) }

    its(:balance) { is_expected.to eq 22.15 }
  end
end
