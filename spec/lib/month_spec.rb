# frozen_string_literal: true

RSpec.describe Month do
  describe '.parse' do
    subject { described_class.parse(params) }

    let(:month) { described_class.new(2023, 5) }

    context 'when it was parsed' do
      let(:params) { '2023-05' }

      it { is_expected.to eq month }
    end

    context do
      before { travel_to Date.new(2023, 5, 1) }

      let(:params) { nil }

      it { is_expected.to eq month }
    end
  end
end
