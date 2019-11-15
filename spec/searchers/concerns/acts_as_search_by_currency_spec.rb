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
      before { expect(relation).to receive(:where).with(currency: 'usd').and_return(:collection) }

      subject { described_class.search relation, currency: 'usd' }

      it { should eq :collection }
    end
  end
end
