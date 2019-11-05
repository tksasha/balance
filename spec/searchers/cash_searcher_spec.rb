# frozen_string_literal: true

RSpec.describe CashSearcher do
  let(:relation) { double }

  describe '#search_by_currency' do
    subject { described_class.search relation, params }

    context do
      let(:params) { { currency: '' } }

      it { should eq relation }
    end

    context do
      let(:params) { { currency: 'usd' } }

      before { expect(relation).to receive(:where).with(currency: 'usd').and_return(:collection) }

      it { should eq :collection }
    end
  end
end
