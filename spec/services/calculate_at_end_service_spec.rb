# frozen_string_literal: true

RSpec.describe CalculateAtEndService do
  subject { described_class.new currency }

  let(:currency) { 'usd' }

  describe '#search_by_currency' do
    context do
      let(:currency) { '' }

      before { allow(Item).to receive(:where).with(currency: 'uah').and_return(:collection) }

      its(:search_by_currency) { should eq :collection }
    end

    context do
      let(:currency) { 'usd' }

      before { allow(Item).to receive(:where).with(currency: 'usd').and_return(:collection) }

      its(:search_by_currency) { should eq :collection }
    end
  end

  describe '#income' do
    before do
      #
      # subject.search_by_currency.income.sum(:sum) -> 15
      #
      allow(subject).to receive(:search_by_currency) do
        double.tap do |a|
          allow(a).to receive(:income) do
            double.tap do |b|
              allow(b).to receive(:sum).with(:sum).and_return(15)
            end
          end
        end
      end
    end

    its(:income) { should eq 15 }
  end

  describe '#expense' do
    before do
      #
      # subject.search_by_currency.expense.sum(:sum) -> 16
      #
      allow(subject).to receive(:search_by_currency) do
        double.tap do |a|
          allow(a).to receive(:expense) do
            double.tap do |b|
              allow(b).to receive(:sum).with(:sum).and_return(16)
            end
          end
        end
      end
    end

    its(:expense) { should eq 16 }
  end

  describe '#call' do
    before { allow(subject).to receive(:income).and_return(10) }

    before { allow(subject).to receive(:expense).and_return(6.5) }

    its(:call) { should eq 3.5 }
  end

  describe '.call' do
    subject { described_class.call('uah') }

    before do
      #
      # described_class.new('uah').call -> 19
      #
      allow(described_class).to receive(:new).with('uah') do
        double.tap { |a| allow(a).to receive(:call).and_return(19) }
      end
    end

    it { should eq 19 }
  end
end
