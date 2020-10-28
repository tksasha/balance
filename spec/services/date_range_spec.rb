# frozen_string_literal: true

RSpec.describe DateRange do
  let(:date) { Date.today }

  subject { described_class.new date }

  its(:month) { should eq date.beginning_of_month..date.end_of_month }

  describe '.month' do
    before do
      #
      # described_class.new(date).month -> :month
      #
      allow(described_class).to receive(:new).with(date) do
        double.tap { |a| allow(a).to receive(:month).and_return(:month) }
      end
    end

    subject { described_class.month date }

    it { should eq :month }
  end
end
