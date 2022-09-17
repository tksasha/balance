# frozen_string_literal: true

RSpec.describe Success do
  subject { described_class.new object }

  let(:object) { double }

  it { is_expected.to be_success }

  it { is_expected.not_to be_failure }
end
