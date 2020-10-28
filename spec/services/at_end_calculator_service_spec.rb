# frozen_string_literal: true

RSpec.describe AtEndCalculatorService do
  let(:params) { {} }

  subject { described_class.new params }

  describe '#search_by_currency' do
    context do
      let(:params) { { currency: '' } }

      before { allow(Item).to receive(:where).with(currency: 'uah').and_return(:collection) }

      its(:search_by_currency) { should eq :collection }
    end

    context do
      let(:params) { { currency: 'usd' } }

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

  describe '#calculate' do
    before { allow(subject).to receive(:income).and_return(10) }

    before { allow(subject).to receive(:expense).and_return(6.5) }

    its(:calculate) { should eq 3.5 }
  end

  describe '.calculate' do
    before do
      #
      # described_class.new(currency: 'uah').calculate -> 19
      #
      allow(described_class).to receive(:new).with(currency: 'uah') do
        double.tap { |a| allow(a).to receive(:calculate).and_return(19) }
      end
    end

    subject { described_class.calculate(currency: 'uah') }

    it { should eq 19 }
  end
end
