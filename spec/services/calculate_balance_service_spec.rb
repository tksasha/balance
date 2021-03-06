# frozen_string_literal: true

RSpec.describe CalculateBalanceService do
  let(:params) { { currency: 'uah' } }

  subject { described_class.new params }

  it { should be_an ApplicationService }

  its(:currency) { should eq 'uah' }

  describe '#at_end' do
    before { allow(CalculateAtEndService).to receive(:call).and_return(21.04) }

    its(:at_end) { should eq 21.04 }
  end

  describe '#sum' do
    context do
      before do
        #
        # Cash.where(currency: 'uah').sum(:sum) -> 21.09
        #
        allow(Cash).to receive(:where).with(currency: 'uah') do
          double.tap do |a|
            allow(a).to receive(:sum).with(:sum).and_return(21.09)
          end
        end
      end

      its(:sum) { should eq 21.09 }
    end

    context do
      let(:params) { { currency: 'usd' } }

      before do
        #
        # Cash.where(currency: 'usd').sum(:sum) -> 42.69
        #
        allow(Cash).to receive(:where).with(currency: 'usd') do
          double.tap do |a|
            allow(a).to receive(:sum).with(:sum).and_return(42.69)
          end
        end
      end

      its(:sum) { should eq 42.69 }
    end
  end

  describe '#call' do
    before { allow(subject).to receive(:sum).and_return(99.999) }

    before { allow(subject).to receive(:at_end).and_return(55.555) }

    its(:call) { should eq 44.44 }
  end

  describe '.call' do
    before do
      #
      # described_class.new(currency: 'usd').call -> 28
      #
      allow(described_class).to receive(:new).with(currency: 'usd') do
        double.tap do |a|
          allow(a).to receive(:call).and_return(28)
        end
      end
    end

    subject { described_class.call(currency: 'usd') }

    it { should eq 28 }
  end
end
