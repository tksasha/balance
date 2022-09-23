# frozen_string_literal: true

RSpec.describe DateRange do
  describe '.parse' do
    let(:current_date_range) { Date.new(2022, 10, 1)..Date.new(2022, 10, 31) }

    before { travel_to Date.new(2022, 10, 23) }

    context 'without arguments' do
      subject { described_class.parse }

      it { is_expected.to eq current_date_range }
    end

    context 'when only `year` is provided' do
      subject { described_class.parse(year: 2022) }

      it { is_expected.to eq current_date_range }
    end

    context 'when only `month` is provided' do
      subject { described_class.parse(month: 10) }

      it { is_expected.to eq current_date_range }
    end

    context 'when month is invalid' do
      subject { described_class.parse(year: 2022, month: 13) }

      it { is_expected.to eq current_date_range }
    end

    context 'with correct arguments' do
      subject { described_class.parse(year: 2022, month: 9) }

      let(:date_range) { Date.new(2022, 9, 1)..Date.new(2022, 9, 30) }

      it { is_expected.to eq date_range }
    end

    context 'when `params` are provided' do
      subject { described_class.parse(params) }

      let(:params) { { year: 2022, month: 9, currency: 'eur' } }

      let(:date_range) { Date.new(2022, 9, 1)..Date.new(2022, 9, 30) }

      it { is_expected.to eq date_range }
    end
  end
end
