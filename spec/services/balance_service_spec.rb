# frozen_string_literal: true

RSpec.describe BalanceService do
  let(:params) { {} }

  subject { described_class.new params }

  describe '#at_end' do
    before { expect(AtEndService).to receive(:at_end).and_return(21.04) }

    its(:at_end) { should eq 21.04 }
  end

  describe '#sum' do
    context do
      before do
        #
        # Cash.where(currency: 'uah').sum(:sum) -> 21.09
        #
        expect(Cash).to receive(:where).with(currency: 'uah') do
          double.tap { |a| expect(a).to receive(:sum).with(:sum).and_return(21.09) }
        end
      end

      its(:sum) { should eq 21.09 }
    end

    context do
      let(:params) { { currency: 'usd' } }

      before do
        #
        # Cash.where(currency: 'usd').sum(:sum) -> 21.09
        #
        expect(Cash).to receive(:where).with(currency: 'usd') do
          double.tap { |a| expect(a).to receive(:sum).with(:sum).and_return(21.09) }
        end
      end

      its(:sum) { should eq 21.09 }
    end
  end

  describe '#balance' do
    before { expect(subject).to receive(:sum).and_return(99.999) }

    before { expect(subject).to receive(:at_end).and_return(55.555) }

    its(:balance) { should eq 44.44 }
  end
end
