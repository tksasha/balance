# frozen_string_literal: true

RSpec.describe CurrencyService do
  subject { described_class.new currency }

  describe '#currency' do
    context do
      let(:currency) { 'usd' }

      its(:currency) { should eq 'usd' }
    end

    context do
      let(:currency) { 'USD' }

      its(:currency) { should eq 'usd' }
    end

    context do
      let(:currency) { nil }

      its(:currency) { should eq 'uah' }
    end

    context do
      let(:currency) { 'unsupported currency' }

      its(:currency) { should eq 'uah' }
    end

    context do
      let(:currency) { 'uah' }

      its(:currency) { should eq 'uah' }
    end

    context do
      let(:currency) { 'rub' }

      its(:currency) { should eq 'rub' }
    end
  end

  describe '.currency' do
    before do
      #
      # described_class.new('usd').currency -> :currency
      #
      allow(described_class).to receive(:new).with('usd') do
        double.tap do |a|
          allow(a).to receive(:currency).and_return(:currency)
        end
      end
    end

    subject { described_class.currency('usd') }

    it { should eq :currency }
  end
end
