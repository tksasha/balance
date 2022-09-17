# frozen_string_literal: true

RSpec.describe CalculateFormulaService do
  describe '.call' do
    subject { described_class.call string }

    let(:string) { nil }

    it { is_expected.to eq 0.0 }

    it { is_expected.to be_a BigDecimal }

    context do
      let(:string) { '2+2' }

      it { is_expected.to eq 4.0 }
    end

    context do
      let(:string) { '10-8' }

      it { is_expected.to eq 2.0 }
    end

    context do
      let(:string) { '3*4' }

      it { is_expected.to eq 12.0 }
    end

    context do
      let(:string) { 'string(2)plus(+)string(2)' }

      it { is_expected.to eq 4.0 }
    end

    context do
      let(:string) { '2++3' }

      it { is_expected.to eq 5.0 }
    end

    context do
      let(:string) { '10----8' }

      it { is_expected.to eq 2.0 }
    end

    context do
      let(:string) { '3***5' }

      it { is_expected.to eq 15.0 }
    end

    context do
      let(:string) { '17...42+2...58' }

      it { is_expected.to eq 20.0 }
    end

    context do
      let(:string) { '-12.75*2.0+++' }

      it { is_expected.to eq(-25.5) }
    end

    context do
      let(:string) { '-12.75*2.0----' }

      it { is_expected.to eq(-25.5) }
    end

    context do
      let(:string) { '-12.75*2.0***' }

      it { is_expected.to eq(-25.5) }
    end

    context do
      let(:string) { '-12.75*2.0...' }

      it { is_expected.to eq(-25.5) }
    end
  end
end
