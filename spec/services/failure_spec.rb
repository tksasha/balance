# frozen_string_literal: true

RSpec.describe Failure do
  subject { described_class.new object }

  let(:object) { double }

  it { should_not be_success }

  it { should be_failure }
end
