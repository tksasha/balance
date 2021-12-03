# frozen_string_literal: true

RSpec.describe Result, type: :service do
  subject { described_class.new object }

  let(:object) { double }

  its(:object) { should eq object }
end
