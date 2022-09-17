# frozen_string_literal: true

RSpec.describe ActsAsController do
  subject { described_class.new }

  let(:described_class) do
    Class.new(ApplicationController) do
      include ActsAsController

      def result
        double
      end
    end
  end

  its(:_helper_methods) { is_expected.to include :result }

  it { is_expected.to delegate_method(:resource).to(:result) }

  it { is_expected.to delegate_method(:success?).to(:result) }

  it { is_expected.to delegate_method(:failure?).to(:result) }
end
