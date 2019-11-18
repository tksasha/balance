# frozen_string_literal: true

RSpec.describe AtEndCalculatorService do
  let(:params) { {} }

  subject { described_class.new params }

  describe '#search_by_currency' do
    context do
      let(:params) { { currency: '' } }

      before { expect(Item).to receive(:where).with(currency: 'uah').and_return(:collection) }

      its(:search_by_currency) { should eq :collection }
    end

    context do
      let(:params) { { currency: 'usd' } }

      before { expect(Item).to receive(:where).with(currency: 'usd').and_return(:collection) }

      its(:search_by_currency) { should eq :collection }
    end
  end

  describe '#income' do
    after { subject.send :income }

    it do
      #
      # subject.search_by_currency.income.sum(:sum)
      #
      expect(subject).to receive(:search_by_currency) do
        double.tap do |a|
          expect(a).to receive(:income) do
            double.tap do |b|
              expect(b).to receive(:sum).with(:sum)
            end
          end
        end
      end
    end
  end

  describe '#expense' do
    after { subject.send :expense }

    it do
      #
      # subject.search_by_currency.expense.sum(:sum)
      #
      expect(subject).to receive(:search_by_currency) do
        double.tap do |a|
          expect(a).to receive(:expense) do
            double.tap do |b|
              expect(b).to receive(:sum).with(:sum)
            end
          end
        end
      end
    end
  end

  describe '#calculate' do
    before { expect(subject).to receive(:income).and_return(10) }

    before { expect(subject).to receive(:expense).and_return(6.5) }

    its(:calculate) { should eq 3.5 }
  end

  describe '.calculate' do
    after { described_class.calculate :params }

    it do
      #
      # described_class.new(:params).calculate
      #
      expect(described_class).to receive(:new).with(:params) do
        double.tap { |a| expect(a).to receive(:calculate) }
      end
    end
  end
end
