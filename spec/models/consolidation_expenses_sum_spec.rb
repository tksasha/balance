require 'rails_helper'

RSpec.describe ConsolidationExpensesSum do
  its(:name) { should eq 'Сума витрат' }

  its(:sum) { should eq 0 }

  context do
    before { described_class.sum = 42.69 }

    its(:sum) { should eq 42.69 }

    after { described_class.sum = 0 }
  end
end
