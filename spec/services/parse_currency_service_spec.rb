# frozen_string_literal: true

RSpec.describe ParseCurrencyService do
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

  describe '#call' do
    let(:currency) { double }

    before { allow(subject).to receive(:currency).and_return(currency) }

    its(:call) { should eq currency }
  end

  describe '.call' do
    before do
      #
      # described_class.new('usd').call
      #
      expect(described_class).to receive(:new).with('usd') do
        double.tap do |a|
          expect(a).to receive(:call)
        end
      end
    end

    subject { described_class.call('usd') }

    it { expect { subject }.not_to raise_error }
  end
end
