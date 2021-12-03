# frozen_string_literal: true

RSpec.describe Success do
  subject { described_class.new object }

  let(:object) { double }

  it { should be_success }

  it { should_not be_failure }
end
