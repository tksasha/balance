# frozen_string_literal: true

RSpec.describe DateRange do
  let(:date) { Date.today }

  subject { described_class.new date }

  its(:month) { should eq date.beginning_of_month..date.end_of_month }

  describe '.month' do
    it do
      #
      # described_class.new(date).month
      #
      expect(described_class).to receive(:new).with(date) do
        double.tap { |a| expect(a).to receive(:month) }
      end
    end

    after { described_class.month date }
  end
end
