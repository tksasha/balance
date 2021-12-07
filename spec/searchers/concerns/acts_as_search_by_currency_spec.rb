# frozen_string_literal: true

RSpec.describe ActsAsSearchByCurrency do
  let :described_class do
    Class.new(ApplicationSearcher) do
      include ActsAsSearchByCurrency
    end
  end

  let(:relation) { double }

  describe '#search_by_currency' do
    context do
      subject { described_class.search relation, currency: '' }

      it { should eq relation }
    end

    context do
      subject { described_class.search relation, currency: 'usd' }

      before { allow(relation).to receive(:where).with(currency: 'usd').and_return(:collection) }

      it { should eq :collection }
    end
  end
end
