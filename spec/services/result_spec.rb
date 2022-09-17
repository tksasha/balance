# frozen_string_literal: true

RSpec.describe Result, type: :service do
  subject { described_class.new object }

  let(:object) { double }

  its(:object) { is_expected.to eq object }

  its(:resource) { is_expected.to eq object }
end
