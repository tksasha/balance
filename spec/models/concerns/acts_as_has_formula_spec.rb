# frozen_string_literal: true

RSpec.describe ActsAsHasFormula do
  let(:described_class) do
    Class.new do
      include ActiveModel::Validations
      include ActiveModel::Validations::Callbacks

      prepend ActsAsHasFormula

      attr_accessor :sum, :formula
    end
  end

  describe '#formula=' do
    before { subject.formula = '2+2' }

    its(:sum) { is_expected.to eq 4.0 }

    its(:formula) { is_expected.to eq '2+2' }
  end
end
